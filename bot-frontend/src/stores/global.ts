import { readable, writable } from 'svelte/store';

export const appVersion = readable('', function start(set) {
  if (window.__token.startsWith('{{')) {
    set('0.0.0');
  } else {
    set(window.__token);
  }

  return function stop() {};
});

export const currentView = writable('main');
