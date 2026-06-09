<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { questionApi } from '@/services/api'

const router = useRouter()
const submitting = ref(false)
const error = ref('')

const form = ref({
  question: '',
  choice_1: '',
  choice_2: '',
  choice_3: '',
  choice_4: '',
})

async function submit() {
  if (!form.value.question.trim()) {
    error.value = 'กรุณากรอกคำถาม'
    return
  }
  if (!form.value.choice_1.trim() || !form.value.choice_2.trim() || !form.value.choice_3.trim() || !form.value.choice_4.trim()) {
    error.value = 'กรุณากรอกคำตอบให้ครบ 4 ข้อ'
    return
  }
  try {
    submitting.value = true
    await questionApi.create(form.value)
    router.push('/')
  } catch {
    error.value = 'ไม่สามารถบันทึกข้อมูลได้'
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-3xl mx-auto p-6">
      <div class="bg-green-600 text-white text-center py-3 rounded-t-lg font-bold text-lg">
        IT 08-2
      </div>

      <div class="bg-white shadow rounded-b-lg p-6">
        <form @submit.prevent="submit" class="space-y-4">
          <p v-if="error" class="text-red-500 text-sm">{{ error }}</p>

          <div class="grid grid-cols-[100px_1fr] items-center gap-3">
            <label class="text-sm text-gray-700 text-right">คำถาม <span class="text-red-500">*</span></label>
            <input
              v-model="form.question"
              type="text"
              class="border border-gray-300 rounded px-3 py-1.5 text-sm w-full focus:outline-none focus:ring-2 focus:ring-green-400"
              placeholder="คำถาม"
            />
          </div>

          <div class="grid grid-cols-[100px_1fr] items-center gap-3">
            <label class="text-sm text-gray-700 text-right">คำตอบ 1</label>
            <input
              v-model="form.choice_1"
              type="text"
              class="border border-gray-300 rounded px-3 py-1.5 text-sm w-full focus:outline-none focus:ring-2 focus:ring-green-400"
              placeholder="คำตอบ 1"
            />
          </div>

          <div class="grid grid-cols-[100px_1fr] items-center gap-3">
            <label class="text-sm text-gray-700 text-right">คำตอบ 2</label>
            <input
              v-model="form.choice_2"
              type="text"
              class="border border-gray-300 rounded px-3 py-1.5 text-sm w-full focus:outline-none focus:ring-2 focus:ring-green-400"
              placeholder="คำตอบ 2"
            />
          </div>

          <div class="grid grid-cols-[100px_1fr] items-center gap-3">
            <label class="text-sm text-gray-700 text-right">คำตอบ 3</label>
            <input
              v-model="form.choice_3"
              type="text"
              class="border border-gray-300 rounded px-3 py-1.5 text-sm w-full focus:outline-none focus:ring-2 focus:ring-green-400"
              placeholder="คำตอบ 3"
            />
          </div>

          <div class="grid grid-cols-[100px_1fr] items-center gap-3">
            <label class="text-sm text-gray-700 text-right">คำตอบ 4</label>
            <input
              v-model="form.choice_4"
              type="text"
              class="border border-gray-300 rounded px-3 py-1.5 text-sm w-full focus:outline-none focus:ring-2 focus:ring-green-400"
              placeholder="คำตอบ 4"
            />
          </div>

          <div class="flex justify-center gap-3 pt-2">
            <button
              type="submit"
              :disabled="submitting"
              class="bg-blue-500 hover:bg-blue-600 disabled:opacity-50 text-white text-sm font-medium px-6 py-2 rounded transition-colors"
            >
              บันทึก
            </button>
            <button
              type="button"
              @click="router.push('/')"
              class="bg-red-500 hover:bg-red-600 text-white text-sm font-medium px-6 py-2 rounded transition-colors"
            >
              ยกเลิก
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
