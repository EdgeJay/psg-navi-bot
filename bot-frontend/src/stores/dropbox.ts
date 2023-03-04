import { writable } from 'svelte/store';

interface ListFileRequestsResponse {
  data: FileRequest[];
}

export interface FileRequest {
  id: string;
  title: string;
  desc: string;
  created_on: string;
  formattedCreatedOn: string;
  url: string;
  file_count: number;
  fileCount: number;
}

function createFileRequests() {
  const { subscribe, set } = writable<FileRequest[]>([]);

  return {
    subscribe,
    async fetchAll(csrfToken: string) {
      fetch('/api/dbx-list-file-requests', {
        method: 'GET',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
          'Content-Type': 'application/json',
          'X-PSGNaviBot-Csrf-Token': csrfToken,
        },
      })
      .then((res) => res.json())
      .then((parsed: ListFileRequestsResponse) => {
        set(parsed.data.map((r) => ({
          ...r,
          formattedCreatedOn: new Date(r.created_on).toLocaleString('en-SG'),
          fileCount: r.file_count,
        })))
      });
    }
  };
}

export const fileRequests = createFileRequests();
