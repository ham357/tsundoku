<template>
  <div class="section">
    <div class="columns">
      <div class="column is-6">
        <div class="box">
          <div class="field">
            <label class="label">Name</label>
            <div class="control">
              <input
                v-model="name"
                class="input is-info"
                type="text"
                placeholder="Name"
              >
            </div>
          </div>

          <div class="field">
            <label class="label">Detail</label>
            <div class="control">
              <input
                v-model="detail"
                class="input is-info"
                type="text"
                placeholder="Detail"
              >
            </div>
          </div>

          <div class="field">
            <label class="label">Price</label>
            <div class="control">
              <input
                v-model="price"
                class="input is-info"
                type="number"
                placeholder="Price"
              >
            </div>
          </div>

          <div class="field is-grouped">
            <div class="control">
              <button
                class="button is-link"
                @click="login"
              >
                Submit
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
    data() {
        return {
            email: '',
            password: ''
        };
    },
    methods: {
        login() {
            const user = {
                email: this.email,
                password: this.password
            };
            this.$store.dispatch('login', user);
        }
    }
};
</script>
