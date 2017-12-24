const vm = new Vue({
  el: '#app',
  data: {
    sleeps: []
  },
  mounted() {
	axios.get("http://localhost:3000/hypnos")
	.then(response => {
		this.sleeps = response.data
	})
  }
});
