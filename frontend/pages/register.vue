<template>
  <v-app id="login" class="primary">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4 lg4>
            <v-card class="elevation-1 pa-3">
              <v-card-text>
                <div class="layout column align-center">
                  <span class="material-icons md-48"> construction </span>
                  <h1 class="flex my-4 primary--text">Developer Tool System</h1>
                </div>
                <!-- <Notification :message="error" v-if="error"/> -->
                <v-form method="post" @submit.prevent="register">
                  <v-text-field
                    append-icon="person"
                    name="username"
                    label="Username"
                    type="text"
                    v-model="model.username"
                    required
                  ></v-text-field>
                  <v-text-field
                    append-icon="person"
                    name="email"
                    label="Email"
                    type="text"
                    v-model="model.email"
                    required
                  ></v-text-field>
                  <v-text-field
                    append-icon="lock"
                    name="password"
                    label="Password"
                    id="password"
                    type="password"
                    v-model="model.password"
                    required
                  ></v-text-field>
                  <v-btn block type="submit" color="primary" :loading="loading"
                  >Register</v-btn
                >
                </v-form>
                <!-- <button  class="title has-text-centered">Register!</button> -->
              </v-card-text>
              <v-card-actions>
                <v-btn icon>
                  <v-icon color="blue">fa fa-facebook-square fa-lg</v-icon>
                </v-btn>
                <v-btn icon>
                  <v-icon color="red">fa fa-google fa-lg</v-icon>
                </v-btn>
                <v-btn icon>
                  <v-icon color="light-blue">fa fa-twitter fa-lg</v-icon>
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  layout: "default",
  data: () => ({
    loading: false,
    model: {
      username: "Phan Trọng",
      email: "admin@example.com",
      password: "password",
    },
  }),

  methods: {
    async login() {
      this.loading = true;
      let data  = await this.$axios.post("/login", this.model);
      if(data.status == 200) {
        this.$router.push('/dashboard');
      }
      console.log(data);
      // this.loading = true;
      // setTimeout(() => {
      //   
      // }, 1000);
    },
    async register() {
      try {
        await this.$axios.post('/register', this.model)

        // await this.$auth.loginWith('local', {
        //   data: {
        //   email: this.email,
        //   password: this.password
        //   },
        // })

        // this.$router.push('/')
      } catch (e) {
        this.error = e.response.data.message
      }
    }
  },
};
</script>
<style scoped lang="css">
.md-48 {
  font-size: 48px;
}
#login {
  height: 50%;
  width: 100%;
  position: absolute;
  top: 0;
  left: 0;
  content: "";
  z-index: 0;
}
</style>
