<script lang="ts">
  import { onMount } from 'svelte';
  import { Router, Route } from "svelte-routing";
  import Home from './routes/Home.svelte';
  import Dropbox from './routes/Dropbox.svelte';
  import DropboxAddFileRequest from './routes/DropboxAddFileRequest.svelte';
  import DropboxListFileRequests from './routes/DropboxListFileRequests.svelte';
  import Help from './routes/Help.svelte';
  import Footer from './lib/Footer.svelte';

  interface InitParams {
    version?: string;
    hash?: string;
  }

  let busy = true

  onMount(() => {
    console.log('app mounted');

    // Init session
    fetch('/api/init-menu-session', {
      method: 'POST',
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        init_data: window.Telegram.WebApp.initData,
      })
    })
    .then((res) => res.json())
    .then((data) => {
      if (data.status !== 'ok') {
        window.Telegram.WebApp.showAlert('Invalid session', () => {
          window.Telegram.WebApp.close();
        });
      } else {
        busy = false
      }
    })
    .catch(() => {
      console.log('Unable to start session');
      window.Telegram.WebApp.showAlert('Unable to start session', () => {
        window.Telegram.WebApp.close();
      });
    });

    // replace "i" elements with data-feather with svg icons
    window.feather.replace();
  });
</script>

<main class="container">
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