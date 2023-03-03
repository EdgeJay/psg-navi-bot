import { writable } from 'svelte/store';

interface AppInfo {
  val: string;
  ver: string;
}

function createAppInfo() {
  const { subscribe, set } = writable<AppInfo>();

  return {
    subscribe,
    async fetch() {
      const data = await window.cookieStore.get('cs');
      const appInfo = JSON.parse(decodeURIComponent(data.value));
      set(appInfo);
    }
  };
}

export const appInfo = createAppInfo();