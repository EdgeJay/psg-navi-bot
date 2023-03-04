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
      if (data) {
        const appInfo = JSON.parse(decodeURIComponent(data.value)) as AppInfo;
        set(appInfo);
      } else {
        set({ val: '', ver: '' });
      }
    }
  };
}

export const appInfo = createAppInfo();