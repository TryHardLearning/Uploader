import { headerComponent } from "./components/header.js";
import { footerComponent } from "./components/footer.js";
import { sidebarComponent } from "./components/sidebar.js";


window.onload = function () {
  const appHeader = document.getElementById("appHeader");
  appHeader.innerHTML = headerComponent();

  const appSidebar = document.getElementById("appSidebar");
  appSidebar.innerHTML = sidebarComponent;


  const appFooter = document.getElementById("footer");
  console.log(appFooter);
  appFooter.innerHTML = footerComponent;
};
