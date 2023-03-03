import { readable } from 'svelte/store';

interface AppInfo {
  val: string;
  ver: string;
}

export const appInfo = readable<AppInfo>(undefined, function start(set) {
  window.cookieStore.get('cs').then((data) => {
    const appInfo = JSON.parse(decodeURIComponent(data.value));
    set(appInfo);
  });

  return function stop() {};
});
