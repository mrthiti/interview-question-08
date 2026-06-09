import { createRouter, createWebHistory } from 'vue-router'
import QuestionList from '@/views/QuestionList.vue'
import QuestionForm from '@/views/QuestionForm.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'list', component: QuestionList },
    { path: '/add', name: 'add', component: QuestionForm },
  ],
})

export default router
