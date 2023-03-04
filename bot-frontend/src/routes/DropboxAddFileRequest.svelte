<script lang="ts">
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import NavBar from "../lib/NavBar.svelte";
  import { appInfo } from "../stores/global";

  const maxChars = 200;

  let busy = false;
  let charsLeft = maxChars;
  let title = '';
  let desc = '';

  $: charsLeft = maxChars - desc.length;

  let csrfToken = '';
  onMount(async () => {
    await appInfo.fetch();
    csrfToken = $appInfo.val;
  });

  function handleSubmit() {
    busy = true;

    fetch('/api/dbx-add-file-request', {
      method: 'POST',
      cache: 'no-cache',
      credentials: 'same-origin',
      headers: {
        'Content-Type': 'application/json',
        'X-PSGNaviBot-Csrf-Token': csrfToken,
      },
      body: JSON.stringify({
        title,
        desc,
      })
    })
    .then((res) => res.json())
    .then((data) => {
      if (data.status !== 'ok') {
        throw new Error('Add new file request failed');
      }
      
      window.Telegram.WebApp.showAlert('File request created', () => {
        navigate('dropbox', { replace: true });
      });
    })
    .catch(() => {
      window.Telegram.WebApp.showAlert('Unable to add new file request', () => {
        busy = false;
      });
    });
  }

  function handleCancel() {
    history.back();
  }
</script>

<section>
  <NavBar title="Add File Request" />
  <form on:submit|preventDefault={handleSubmit}>
      <label for="txt-filerequest-title">
          File Request Title
          <input type="text" bind:value={title} placeholder="Enter title of file request here" required />
      </label>
      <label for="txt-filerequest-desc">
          File Request Description
          <textarea bind:value={desc} placeholder="Maximum 200 characters" maxlength="200" rows="5"></textarea>
      </label>
      <p class="chars-left">Characters left: {charsLeft}</p>
      <button type="submit">Create</button>
      <button on:click={handleCancel} class="secondary">Cancel</button>
  </form>
  {#if busy}
    <dialog open>
      <article>
        <header>
          Creating Dropbox File Request
        </header>
        <p aria-busy="true">Please wait...</p>
      </article>
    </dialog>
  {/if}
</section>

<style>
  .chars-left {
    font-size: 0.9rem;
    text-align: right;
  }
</style>