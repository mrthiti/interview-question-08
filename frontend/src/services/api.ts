import axios from 'axios'

const http = axios.create({ baseURL: 'http://localhost:8080/api' })

export interface Question {
  ID: number
  order_number: number
  question: string
  choice_1: string
  choice_2: string
  choice_3: string
  choice_4: string
}

export interface CreateQuestionPayload {
  question: string
  choice_1: string
  choice_2: string
  choice_3: string
  choice_4: string
}

export const questionApi = {
  getAll: () => http.get<Question[]>('/questions'),
  create: (payload: CreateQuestionPayload) => http.post<Question>('/questions', payload),
  remove: (id: number) => http.delete(`/questions/${id}`),
}
