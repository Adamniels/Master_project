import { createRouter, createWebHistory } from "vue-router";
import LoginPage from "../views/LoginPage.vue";
import HomePage from "../views/HomePage.vue";

const routes = [
  {
    path: "/",
    redirect: "/loginpage",
  },
  {
    path: "/loginpage",
    component: LoginPage,
  },
  {
    path: "/homepage",
    component: HomePage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");

  if (!token && to.path !== "/loginpage") {
    // Om ingen token finns och man försöker gå till en annan sida än /login
    next("/loginpage");
  } else {
    // Allt annat är okej
    next();
  }
});

export default router;
