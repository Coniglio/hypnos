const vm = new Vue({
  el: '#app',
  data: {
    sleeps: []
  },
  mounted() {
	axios.get("http://tk2-208-13803.vs.sakura.ne.jp:3000/hypnos")
	.then(response => {
		this.sleeps = response.data
	})
  }
});