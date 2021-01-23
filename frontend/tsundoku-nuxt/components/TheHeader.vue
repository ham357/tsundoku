<template>
  <nav class="navbar">
    <div class="navbar-brand">
      <nuxt-link class="navbar-item subtitle is-4" to="/">
        Tsundoku
      </nuxt-link>
    </div>

    <div class="navbar-menu">
      <div class="navbar-start">
        <div class="navbar-item">
          <div class="field has-addons" />
          <div
            v-if="isLoggedIn"
          >
            <nuxt-link class="navbar-item subtitle is-5" to="/books">
              本を積む
            </nuxt-link>
          </div>
        </div>
      </div>

      <div class="navbar-end">
        <div class="navbar-item">
          <div
            v-if="isLoggedIn"
            class="buttons"
          >
            <a
              href="#"
              class="button is-ligth"
              @click.prevent="logout"
            >
              Log out
            </a>
          </div>
          <div
            v-else
            class="buttons"
          >
            <nuxt-link
              class="button is-primary"
              to="/signup"
            >
              <strong>Sign up</strong>
            </nuxt-link>
            <nuxt-link
              class="button is-light"
              to="/login"
            >
              Log in
            </nuxt-link>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
    computed: {
        isLoggedIn () {
            return !!this.$store.getters.isLoggedIn;
        }
    },
    methods: {
        logout () {
            this.$store.dispatch('logout');
        }
    }
};
</script>
