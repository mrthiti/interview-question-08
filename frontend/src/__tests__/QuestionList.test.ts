import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount, flushPromises } from '@vue/test-utils'
import { createRouter, createMemoryHistory } from 'vue-router'
import QuestionList from '@/views/QuestionList.vue'
import * as api from '@/services/api'

vi.mock('@/services/api')

const mockQuestions = [
  { ID: 1, order_number: 1, question: 'Q1', choice_1: '', choice_2: '', choice_3: '', choice_4: '' },
  { ID: 2, order_number: 2, question: 'Q2', choice_1: 'A', choice_2: '', choice_3: '', choice_4: '' },
]

function makeRouter() {
  return createRouter({
    history: createMemoryHistory(),
    routes: [
      { path: '/', component: QuestionList },
      { path: '/add', component: { template: '<div/>' } },
    ],
  })
}

describe('QuestionList', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    vi.mocked(api.questionApi.getAll).mockResolvedValue({ data: mockQuestions } as any)
    vi.mocked(api.questionApi.remove).mockResolvedValue({ data: {} } as any)
  })

  it('renders the IT 08-1 header', async () => {
    const wrapper = mount(QuestionList, { global: { plugins: [makeRouter()] } })
    await flushPromises()
    expect(wrapper.text()).toContain('IT 08-1')
  })

  it('displays questions fetched from API', async () => {
    const wrapper = mount(QuestionList, { global: { plugins: [makeRouter()] } })
    await flushPromises()
    expect(wrapper.text()).toContain('Q1')
    expect(wrapper.text()).toContain('Q2')
  })

  it('shows empty state when no questions', async () => {
    vi.mocked(api.questionApi.getAll).mockResolvedValue({ data: [] } as any)
    const wrapper = mount(QuestionList, { global: { plugins: [makeRouter()] } })
    await flushPromises()
    expect(wrapper.text()).toContain('ยังไม่มีข้อสอบ')
  })

  it('calls delete API and refreshes when delete button clicked', async () => {
    const wrapper = mount(QuestionList, { global: { plugins: [makeRouter()] } })
    await flushPromises()

    const deleteButtons = wrapper.findAll('button').filter(b => b.text() === 'ลบ')
    await deleteButtons[0]!.trigger('click')
    await flushPromises()

    expect(api.questionApi.remove).toHaveBeenCalledWith(1)
    expect(api.questionApi.getAll).toHaveBeenCalledTimes(2)
  })

  it('shows error message when API fails', async () => {
    vi.mocked(api.questionApi.getAll).mockRejectedValue(new Error('Network error'))
    const wrapper = mount(QuestionList, { global: { plugins: [makeRouter()] } })
    await flushPromises()
    expect(wrapper.text()).toContain('ไม่สามารถโหลดข้อมูลได้')
  })
})
