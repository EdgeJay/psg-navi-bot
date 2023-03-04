<script lang="ts">
  import { onMount } from 'svelte';
  import { Router, Route } from "svelte-routing";
  import Home from './routes/Home.svelte';
  import Dropbox from './routes/Dropbox.svelte';
  import DropboxAddFileRequest from './routes/DropboxAddFileRequest.svelte';
  import DropboxListFileRequests from './routes/DropboxListFileRequests.svelte';
  import Help from './routes/Help.svelte';
  import Footer from './lib/Footer.svelte';
  import { appInfo } from './stores/global';

  interface InitResponse {
    status: string;
    ver: string;
  }

  let busy = true;

  let colorScheme = window.Telegram.WebApp.colorScheme;

  onMount(async () => {
    console.log('app mounted');

    // Init session
    try {
      const res = await fetch('/api/init-menu-session', {
        method: 'POST',
        cache: 'no-cache',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          init_data: window.Telegram.WebApp.initData,
        })
      });
      const data = (await res.json()) as InitResponse;
      if (data.status !== 'ok') {
        window.Telegram.WebApp.showAlert('Invalid session', () => {
          window.Telegram.WebApp.close();
        });
      } else {
        busy = false;

        // force app version shown in UI to be refreshed
        appInfo.fetch();
      }
    } catch (err: unknown) {
      window.Telegram.WebApp.showAlert('Unable to start session', () => {
        window.Telegram.WebApp.close();
      });
    }
  });
</script>

<main class="container" data-theme={colorScheme}>
  <!-- Verify app version, init session first before showing rest of app -->
  {#if !busy}
    <Router>
      <Route path="dropbox"><Dropbox /></Route>
      <Route path="dbx-add-file-request"><DropboxAddFileRequest /></Route>
      <Route path="dbx-list-file-requests"><DropboxListFileRequests /></Route>
      <Route path="help"><Help /></Route>
      <Route path="/"><Home /></Route>
    </Router>
  {/if}
  <Footer />
</main>

<style>

</style>