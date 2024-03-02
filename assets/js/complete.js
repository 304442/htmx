document.addEventListener("DOMContentLoaded", async ()=>{
    // fetch publishable key and init stripe
    const { publishableKey } = await fetch("/config").then(r => r.json())
    const stripe = Stripe(publishableKey)


    // use prarams to retrieve pmeny intet

    const params = new URLSearchParams(window.location.href)

    const clientSecret = params.get("payment_intent_client_secret")
    

    const {paymentIntent} = await stripe.retrievePaymentIntent(clientSecret)
    // render the payment intent

    const per = document.getElementById("payment-intent")

    per.innerText = JSON.stringify(paymentIntent, null, 2)
})