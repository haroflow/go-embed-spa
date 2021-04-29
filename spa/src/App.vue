<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">

	<br>
	<button @click="addRandomItem">Add random item</button>

	<h4>Items from Go:</h4>
	<p v-if="error" class="error">Could not load items from the backend</p>
	<div>
		<p v-for="item, idx in items" :key="item">
			{{ idx }} - {{ item }}
		</p>
	</div>
  </div>
</template>

<script>
export default {
  name: 'App',

	data: function () {
		return {
			error: null,
			items: []
		}
	},

	mounted() {
		this.loadItems()
	},

	methods: {
		// Load items from the Go backend
		loadItems() {
			fetch("/items")
				.then(res => res.json())
				.then(res => this.items = res)
				.catch(err => this.error = err)
		},

		// Ask the backend to add a random item to the list
		async addRandomItem() {
			// Add item
			let res = await fetch("/items", {
				method: "POST",
			}).then(res => res.json())

			console.log("Go added a new item:", res.new_item)
			
			// Reload list
			this.loadItems()
		}
	},
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.error {
	color: red;
}
</style>
