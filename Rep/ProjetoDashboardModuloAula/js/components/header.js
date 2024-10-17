export const headerComponent = function(){
   return `
   <nav class="navbar navbar-light bg-light fixed-top menu-top">
   <div class="d-flex justify-content-between col-12 col-md-3 col-lg-2">
       <a href="index.html" class="navbar-brand">Dashboard</a>
       <button type="button" class="navbar-toggler d-md-none d-lg-none collapsed" data-bs-toggle="collapse"
           data-bs-target="#sidebar">
           <span class="navbar-toggler-icon"></span>
       </button>
   </div>

   <div class="col-12 col-md-4 col-lg-2 ps-3">
       <input type="search" name="search" id="search" placeholder="Procurar..." class="form-control">
   </div>

   <div class="d-flex justify-content-end pe-5 col-12 col-md-5 col-lg-8">
       <div class="dropdown">
           <button class="avatar border-0 rounded-circle bg-light dropdown-toggle" data-bs-toggle="dropdown">
               <img src="img/avatar5.png" alt="Avatar" title="Avatar">
           </button>
           <ul class="dropdown-menu">
               <li><a href="#" class="dropdown-item">Configurações</a></li>
               <li><a href="#" class="dropdown-item">Mensagens</a></li>
               <li><a href="#" class="dropdown-item">Sair</a></li>
           </ul>
       </div>

   </div>
</nav>
   `;
};
