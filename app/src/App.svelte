<script>
import { onMount } from "svelte"
let todos = []
let thisTodo = ''
let thisUrl = window.location.protocol + '//' + window.location.host;

onMount(async () => {
  const res = await fetch(thisUrl + '/api/crud');
  todos = await res.json();
});

function addTodo(e) {
  e.preventDefault();

  fetch(thisUrl + '/api/crud', {
    method: 'POST',
    body: JSON.stringify({
      todo: thisTodo
    })
  }).then(
      todos = [thisTodo, ...todos],
      thisTodo = ''
  ).catch(err => {alert(err)})
}

function clearTodos(e) {
  e.preventDefault();
  fetch(thisUrl + '/api/crud', {
    method: 'DELETE',
  }).then(
      todos = [],
      thisTodo = ''
  ).catch(err => {alert(err)})
}
</script>

<main>
  <h1>todos</h1>
  <form on:submit={addTodo}>
    <input bind:value={thisTodo} placeholder="enter a todo" />
  </form>
  <div>
    <button on:click={addTodo}>Submit</button>
    <button on:click={clearTodos}>Clear</button>
  </div>
  {#each todos as todo}
    <p>{todo}</p>
  {/each}
</main>

<style>
main {
  text-align: center;
  padding: 1em;
  max-width: 240px;
  margin: 0 auto;
}

h1 {
  color: #ff3e00;
  text-transform: uppercase;
  font-size: 4em;
  font-weight: 100;
}

@media (min-width: 640px) {
  main {
    max-width: none;
  }
}
</style>
