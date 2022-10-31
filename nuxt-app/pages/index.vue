<template>
<div>
  <input v-model="loginRequest.username">
  <input v-model="loginRequest.password">
  <button @click="login">login</button>
  {{ loginResponse.token }}
</div>
</template>

<script setup lang="ts">
import axios from "~/service/axios";
import {authUrl} from "~/service/proto_gen/auth/auth_url";
import {$ref} from "vue/macros";
import {auth} from "~/service/proto_gen/auth/auth_pb";

let loginRequest = $ref({} as auth.LoginRequest)
let loginResponse = $ref({} as auth.LoginResponse)
const login = () => {
  axios.post(authUrl.Login, loginRequest).then(rep => {
    loginResponse = auth.LoginResponse.fromObject(rep.data)
  })
}
</script>

<style scoped>

</style>