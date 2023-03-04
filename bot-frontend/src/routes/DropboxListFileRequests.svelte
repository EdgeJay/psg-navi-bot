<script lang="ts">
  import { onMount } from 'svelte';
  import NavBar from "../lib/NavBar.svelte";
  import { appInfo } from "../stores/global";
  import { fileRequests } from "../stores/dropbox";

  let fetchFileRequestsPromise = Promise.resolve();
  
  onMount(async () => {
    await appInfo.fetch();
    const csrfToken = $appInfo.val;
    fetchFileRequestsPromise = fileRequests.fetchAll(csrfToken);
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
        {#await fetchFileRequestsPromise}
          <tr></tr>
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
      </tbody>
  </table>
</section>