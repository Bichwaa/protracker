import { createRouter, createWebHistory } from 'vue-router'
import Home from '../components/Home.vue';
import ProjectDetails from '../components/ProjectDetails.vue';
import Notes from '../components/Notes.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/project-details/:id',
    name: 'Details',
    component: ProjectDetails
  },
  {
    path:"/notes",
    name:"Notes",
    component: Notes
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
