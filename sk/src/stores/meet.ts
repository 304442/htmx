import { writable } from 'svelte/store'

interface Meet {
  duration: number // minute
  date: Date | null
  zoneTime: string
  start: string
  // end: string
  name: string
  email: string
  note: string
}

export const meet = writable<Meet>({
  duration: 15,
  date: null,
  start: '',
  zoneTime: '',
  name: '',
  email: '',
  note: '',
})

export const updateMeet = (newMeet: Partial<Meet>) => {
  meet.update((existingMeet) => ({
    ...existingMeet,
    ...newMeet,
  }))
}

meet.subscribe((meet) => {
  console.table({ meet })
})