interface Cookie {
  domain: string;
  expires: number;
  name: string;
  path: string;
  sameSite: string;
  secure: boolean;
  value: string;
}

interface Window {
  __token: string;
  __version: string;
  feather: {
    replace: () => void;
  };
  cookieStore: {
    get: (name: string) => Promise<Cookie>
  },
  Telegram: {
    WebApp: {
      showAlert: (message: string, callback: () => void) => void,
      openLink: (url: string) => void,
    },
  },
}