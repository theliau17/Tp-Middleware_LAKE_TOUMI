<script>
    import {resources, resourcesError, getResources, postResources, putResource, deleteResource} from "$lib/stores/resources.js";
    import DeleteModal from "../../components/delete-modal.svelte"
    import {onMount} from "svelte";
    import {SvelteMap} from "svelte/reactivity";

    let newUcaId = $state("")
    let newUcaName = $state("")
    let editable = new SvelteMap()

    let deleteModal = $state();


    onMount(() => {
        getResources()
    })

    function addResource() {
        console.log(`Adding resource with name ${newUcaName} and ucaId ${newUcaId}`)
        postResources(newUcaId, newUcaName)
            .then(() => {
                newUcaName = ""
                newUcaId = ""
            })
            .catch(() => {
            })
    }

    function deleteResourceById(id) {
        console.log(`Deleting resource ${id}`)
        deleteResource(id)
            .then(() => {
                editable.set(id, false)
            })
            .catch(() => {
            })
    }

    function editResource(id, name, ucaId) {
        console.log(`Editing resource ${id} with name ${name} and ucaId ${ucaId}`)
        putResource(id, ucaId, name)
            .then(() => {
                editable.set(id, false)
            })
            .catch(() => {
            })
    }
</script>

<div class="d-flex flex-column">
    <h1 class="text-center color-yellow mb-lg-5">Resources</h1>
    <div class="d-flex flex-column mb-5">
        <h5 class="text-center fw-semibold mb-2">Add a resource : </h5>
        <form class="d-flex justify-content-center" onsubmit="{addResource}">
            <div class="d-flex me-4">
                <label for="id" class="fs-5 align-self-center me-2">UCA ID</label>
                <input id="id" class="form-control-sm" placeholder="56529" bind:value={newUcaId}/>
            </div>
            <div class="d-flex me-4">
                <label for="name" class="fs-5 align-self-center me-2">Name</label>
                <input id="name" class="form-control-sm" placeholder="M1 - Tutorat L2" bind:value={newUcaName}/>
            </div>
            <input class="btn btn-yellow" type="submit" value="Submit"/>
        </form>
    </div>
    {#if $resourcesError}
        <div class="text-center text-danger form-text fs-5">
            An error occurred : {$resourcesError}
        </div>
    {/if}
    <table class="table w-50 align-self-center">
        <thead>
        <tr>
            <th class="color-yellow" scope="col">Name</th>
            <th class="color-yellow" scope="col">UCA ID</th>
            <th class="color-yellow" scope="col">ID</th>
            <th scope="col">Actions</th>
        </tr>
        </thead>
        <tbody>
        {#each $resources as resource}
            <tr>
                <td>
                    {#if !editable.get(resource.id)}
                        {resource.name}
                    {:else}
                        <input class="form-control-sm" bind:value={resource.name} />
                    {/if}
                </td>
                <td>
                    {#if !editable.get(resource.id)}
                        {resource.ucaId}
                    {:else}
                        <input class="form-control-sm" bind:value={resource.ucaId} />
                    {/if}
                </td>
                <td>{resource.id}</td>
                <td class="justify-content-around">
                    {#if !editable.get(resource.id)}
                        <button class="bg-transparent border-0" onclick="{() => {editable.set(resource.id, true)}}">
                            <span class="fa fa-pen-to-square text-warning"></span>
                        </button>
                        <button class="bg-transparent border-0" onclick={() => { deleteModal.show(resource.id, resource.name) }}>
                            <span class="fa fa-trash-can text-danger"></span>
                        </button>
                    {:else}
                        <button class="bg-transparent border-0" onclick="{() => {editResource(resource.id, resource.name, resource.ucaId)}}">
                            <span class="fa fa-check" style="color: #0b6f33"></span>
                        </button>
                    {/if}
                </td>
            </tr>
        {/each}
        </tbody>
    </table>
    <DeleteModal bind:this={deleteModal} uniqueName="resources" deleteFunction="{deleteResourceById}" />
</div>