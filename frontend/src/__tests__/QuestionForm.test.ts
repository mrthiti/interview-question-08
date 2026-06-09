import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount, flushPromises } from '@vue/test-utils'
import { createRouter, createMemoryHistory } from 'vue-router'
import QuestionForm from '@/views/QuestionForm.vue'
import * as api from '@/services/api'

vi.mock('@/services/api')

function makeRouter() {
  return createRouter({
    history: createMemoryHistory(),
    routes: [
      { path: '/', component: { template: '<div/>' } },
      { path: '/add', component: QuestionForm },
    ],
  })
}

describe('QuestionForm', () => {
  beforeEach(() => {
    vi.mocked(api.questionApi.create).mockResolvedValue({ data: {} } as any)
  })

  it('renders the IT 08-2 header', () => {
    const wrapper = mount(QuestionForm, { global: { plugins: [makeRouter()] } })
    expect(wrapper.text()).toContain('IT 08-2')
  })

  it('shows validation error when submitting without question', async () => {
    const wrapper = mount(QuestionForm, { global: { plugins: [makeRouter()] } })
    await wrapper.find('form').trigger('submit')
    await flushPromises()
    expect(wrapper.text()).toContain('กรุณากรอกคำถาม')
    expect(api.questionApi.create).not.toHaveBeenCalled()
  })

  it('calls create API with form data when submitted', async () => {
    const wrapper = mount(QuestionForm, { global: { plugins: [makeRouter()] } })
    const inputs = wrapper.findAll('input')

    await inputs[0]!.setValue('My Question')
    await inputs[1]!.setValue('Option A')
    await inputs[2]!.setValue('Option B')
    await inputs[3]!.setValue('Option C')
    await inputs[4]!.setValue('Option D')
    await wrapper.find('form').trigger('submit')
    await flushPromises()

    expect(api.questionApi.create).toHaveBeenCalledWith({
      question: 'My Question',
      choice_1: 'Option A',
      choice_2: 'Option B',
      choice_3: 'Option C',
      choice_4: 'Option D',
    })
  })

  it('navigates back on cancel', async () => {
    const router = makeRouter()
    await router.push('/add')
    const wrapper = mount(QuestionForm, { global: { plugins: [router] } })

    const cancelBtn = wrapper.findAll('button').find(b => b.text() === 'ยกเลิก')
    await cancelBtn!.trigger('click')
    await flushPromises()

    expect(router.currentRoute.value.path).toBe('/')
  })

  it('shows error when API call fails', async () => {
    vi.mocked(api.questionApi.create).mockRejectedValue(new Error('Network error'))
    const wrapper = mount(QuestionForm, { global: { plugins: [makeRouter()] } })
    const inputs = wrapper.findAll('input')
    await inputs[0]!.setValue('Test question')
    await wrapper.find('form').trigger('submit')
    await flushPromises()
    expect(wrapper.text()).toContain('กรุณากรอกคำตอบให้ครบ 4 ข้อ')
  })
})
