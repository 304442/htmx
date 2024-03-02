document.addEventListener("DOMContentLoaded", async () => {
    // fetch publishable key and initi stripe

    const { publishableKey } = await fetch("/config").then(r => r.json())
    const stripe = Stripe(publishableKey)

    // fetch PI client secter

    const { client_secret } = await fetch("/cpi", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({})
    }).then(r => r.json())

    // mount the element
    const elements = stripe.elements({clientSecret: client_secret })

    const paymentElement = elements.create('payment')

    paymentElement.mount("#payment-element")

    const form = document.getElementById("payment-form")

    form.addEventListener("submit", async(e)=>{
        e.preventDefault()
        const{error } = await stripe.confirmPayment({
            elements,
            confirmParams: {
                return_url: new URL(window.location.href).origin+"/success"
            }
        })

        console.log(window.location.href)

        if(error){
            console.log("failed to pay")
            const messages = document.getElementById("error-messages")
            messages.innerText = error.message;
        }else{
            console.log("payment succesfull")
        }
    })
})

