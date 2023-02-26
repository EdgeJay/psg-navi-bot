<script lang="ts">
  import NavBar from "../lib/NavBar.svelte";
</script>

<section>
  <NavBar title="List File Requests" />
  <table role="grid">
      <thead>
          <tr>
              <th>#</th>
              <th>Title</th>
              <th>Actions</th>
          </tr>
      </thead>
      <tbody></tbody>
  </table>
</section>