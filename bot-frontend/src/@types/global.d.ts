interface Window {
  __token: string;
  __version: string;
  feather: {
    replace: () => void;
  };
  cookieStore: {
    get: (name: string) => Promise<string>
  },
}