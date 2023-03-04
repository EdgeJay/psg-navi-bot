import { writable } from 'svelte/store';

interface AppInfo {
  val: string;
  ver: string;
}

const blankInfo = { val: '', ver: '' };

function createAppInfo() {
  const { subscribe, set } = writable<AppInfo>(blankInfo);

  return {
    subscribe,
    async fetch() {
      const data = await window.cookieStore.get('cs');
      if (data) {
        const appInfo = JSON.parse(decodeURIComponent(data.value)) as AppInfo;
        set(appInfo);
      } else {
        set(blankInfo);
      }
    }
  };
}

export const appInfo = createAppInfo();