<!-- RequiredLabel.svelte -->
<script lang="ts">
    export let forId: string;
    export let label: string;
    export let required: boolean = false;
</script>

<label for={forId} class="block text-sm font-medium text-gray-700">
    {label}
    {#if required}
        <span class="text-red-500">*</span>
    {:else}
        <span class="text-gray-400 text-sm"></span>
    {/if}
</label>
