<script lang="ts">
  import { onMount } from 'svelte';
  import NavBar from "../lib/NavBar.svelte";
  import { appInfo } from "../stores/global";
  import { fileRequests } from "../stores/dropbox";

  async function fetchFileRequests(): Promise<void> {
    await appInfo.fetch();
    const csrfToken = $appInfo.val;
    return fileRequests.fetchAll(csrfToken);
  }

  let fetchFileRequestsPromise: Promise<void> | undefined;

  onMount(() => {
    fetchFileRequestsPromise = fetchFileRequests();
  });

  function titleClickHandler(url: string) {
    return function handler() {
      window.Telegram.WebApp.openLink(url);
    }
  }
</script>

<section>
  <NavBar title="List File Requests" />
  <table role="grid">
      <thead>
          <tr>
              <th>Title</th>
              <th>Created On</th>
              <th>File Count</th>
              <th>Actions</th>
          </tr>
      </thead>
      <tbody>
        {#if fetchFileRequestsPromise}
          {#await fetchFileRequestsPromise}
            <tr>
              <td colspan="4">
                <p aria-busy>Fetching file requests...</p>
              </td>
            </tr>
          {:then _} 
            {#each $fileRequests as req}
              <tr>
                <td><a href="{req.url}" on:click|preventDefault={titleClickHandler(req.url)}>{req.title}</a></td>
                <td>{req.formattedCreatedOn}</td>
                <td>{req.fileCount}</td>
                <td></td>
              </tr>
            {/each}
          {/await}
        {:else}
          <tr>
            <td colspan="4">
              <p aria-busy>Fetching file requests...</p>
            </td>
          </tr>
        {/if}
      </tbody>
  </table>
</section>