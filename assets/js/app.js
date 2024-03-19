const modalToggles = document.querySelectorAll('.modal-toggle');
const modals = document.querySelectorAll('.modal');
const closeModals = document.querySelectorAll('.close-modal');
const modal = document.querySelector('.modal');


modalToggles.forEach(toggle => {
    toggle.addEventListener('click', () => {
      const modalId = toggle.dataset.modal;
      const modal = document.getElementById(modalId);
      modals.forEach(m => {
        if (m !== modal) {
          m.classList.add('hidden');
        }
      });
      modal.classList.toggle('hidden');
    });
  });
  
  closeModals.forEach((c)=>c.addEventListener('click', () => {
    const modal = c.closest('.modal');
    modal.classList.add('hidden');
  }))
  
  window.addEventListener('click', (event) => {
    if (!event.target.classList.contains('modal-toggle')) {

      if (!event.target.closest('.modal')) {
        modals.forEach(modal => {
          modal.classList.add('hidden');
        });
      }
    }
  });

let showABout = false
function toggleAbout(){
    showABout = !showABout

    if (showABout){
        document.querySelector("#about_section").classList.remove("hidden")
        document.getElementById("toggleAboutIcon").innerText = "arrow_drop_up"
        
    }else{
        document.querySelector("#about_section").classList.add("hidden")
        document.getElementById("toggleAboutIcon").innerText = "arrow_drop_down"
    }
    
}