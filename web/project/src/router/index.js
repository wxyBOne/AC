/*
 * @Author: wxyBone 13638456960@163.com
 * @Date: 2024-06-13 11:11:43
 * @LastEditors: wxyBone 13638456960@163.com
 * @LastEditTime: 2024-12-19 18:49:08
 * @FilePath: \dapp_vue\project\src\router\index.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Animation",
    component: () => import("@/views/AnimationPage.vue"), 
  },
  {
    path: "/Home",
    name: "Home",
    component: () => import("@/views/Home.vue"), 
    children: [
      {
        path: "/Shop",
        name: "Shop",
        component: () => import("@/components/Shop.vue"), 
      },
      {
        path: "/Product/:id",
        name: "Product",
        component: () => import("@/components/Product.vue"), 
      },
      {
        path: "/Cart",
        name: "Cart",
        component: () => import("@/components/Cart.vue"), 
      },
      {
        path: "/TransBlock",
        name: "TransBlock",
        component: () => import("@/views/TransBlock.vue"), 
        children: [
          {
            path: "/Published",
            name: "Published",
            component: () => import("@/components/Published.vue") 
          },
          {
            path: "/Purchased",
            name: "Purchased",
            component: () => import("@/components/Purchased.vue") 
          },
          {
            path: "/Sold",
            name: "Sold",
            component: () => import("@/components/Sold.vue") 
          }
        ]
      }
    ],
    redirect: "/Shop"
  },
  {
    path: "/Login",
    name: "Login",
    component: () => import("@/views/Login.vue"), 
  },
  {
    path: "/Publish",
    name: "Publish",
    component: () => import("@/views/Publish.vue"), 
  },
  {
    path: "/ProductItem",
    name: "ProductItem",
    component: () => import("@/components/ProductItem.vue"), 
  },
  {
    path: "/B",
    name: "B",
    component: () => import("@/views/B.vue"), 
    children: [
      {
        path: "/UserAdmin", 
        component: () => import("@/components/UserAdmin.vue"), 
      },
      {
        path: '/Overview', 
        component: () => import("@/components/Overview.vue") 
      },
      {
        path: '/ArtCheck', 
        component: () => import("@/components/ArtCheck.vue") 
      },
    ],
    redirect: "/Overview"
  },
  {
    path: "/AdminLog",
    name: "AdminLog",
    component: () => import("@/views/AdminLog.vue"), 
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
