import { writable } from 'svelte/store';

export const editorDataStore = writable({
	consoleValue: null
})