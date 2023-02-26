<script lang="ts">
  import NavBar from "../lib/NavBar.svelte";

  function handleCancel() {
    history.back();
  }
</script>

<section>
  <NavBar title="Add File Request" />
  <form id="form-add-dropbox-file-request">
      <label for="txt-filerequest-title">
          File Request Title
          <input type="text" id="txt-filerequest-title" name="txt-filerequest-title" required />
      </label>
      <label for="txt-filerequest-desc">
          File Request Description
          <textarea id="txt-filerequest-desc" name="txt-filerequest-desc" maxlength="200" rows="5"></textarea>
      </label>
      <button type="submit">Create</button>
      <button on:click={handleCancel} class="secondary">Cancel</button>
  </form>
</section>