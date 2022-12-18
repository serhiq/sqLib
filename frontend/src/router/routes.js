import ImagesPage from "pages/ImagesPage";
import EditImagePage from "pages/EditImagePage";
import UploadPage from "pages/UploadPage";
import ErrorNotFound from "pages/ErrorNotFound";
import RegisterUserView from "components/RegisterUserView";
import LoginUserView from "components/LoginUserView";
import ImageDisplay from "pages/ImagePage";

import router from "src/router/index";

const routes = [
  {
    name: "root",
    path: '/',
    component: ImagesPage,
  },
  // {
  //   path: '/upload',
  //   component: UploadPage,
  // },
  {
    path: '/upload',
    component: () => import('pages/UploadPage'),
    meta: { requiresAuth: true }
  },
  {
    path: '/edit',
    component: () => import('pages/EditImagePage'),
    meta: { requiresAuth: true }
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  },
  {
    path: '/signup',
    name: 'singup',
    component: RegisterUserView
  },
  {
    path: '/login',
    name: 'authenticate',
    component: LoginUserView
  },
  {
    path: '/images/:id',
    name: 'ImageDisplay',
    component: ImageDisplay
  },
]


export default routes
