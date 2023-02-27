<script lang="ts">
  import MenuButton from "../lib/MenuButton.svelte";
</script>

<section>
  <MenuButton to="dropbox" iconType="devicons" icon="dropbox">
    <span slot="label">Dropbox</span>
  </MenuButton>
  <MenuButton to="help" iconType="feather" icon="help-circle">
    <span slot="label">Ask me a question</span>
  </MenuButton>
</section>