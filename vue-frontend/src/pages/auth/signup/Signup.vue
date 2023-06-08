<template>
  <form @submit.prevent="submit">
    <va-input
      v-model="v$.form.firstName.$model"
      class="mb-3"
      type="text"
      label="Nama Depan"
      :error="!!v$.form.firstName.$errors.length"
      :error-messages="v$.form.firstName.$errors.length > 0 ? v$.form.firstName.$errors[0].$message:''"
    />
    <va-input
      v-model="v$.form.lastName.$model"
      class="mb-3"
      type="text"
      label="Nama Belakang"
      :error="!!v$.form.lastName.$errors.length"
      :error-messages="v$.form.lastName.$errors.length > 0 ? v$.form.lastName.$errors[0].$message:''"
    />
    <va-input
      v-model="v$.form.email.$model"
      class="mb-3"
      type="email"
      label="Email"
      :error="!!v$.form.email.$errors.length"
      :error-messages="v$.form.email.$errors.length > 0 ? v$.form.email.$errors[0].$message:''"
    />

    <va-input
      v-model="v$.form.password.$model"
      class="mb-3"
      type="password"
      label="Password"
      :error="!!v$.form.password.$errors.length"
      :error-messages="v$.form.password.$errors.length > 0 ? v$.form.password.$errors[0].$message:''"
    />

    <va-input
      v-model="v$.form.passwordConfirm.$model"
      class="mb-3"
      type="password"
      label="Password Confirmation"
      :error="!!v$.form.passwordConfirm.$errors.length"
      :error-messages="v$.form.passwordConfirm.$errors.length > 0 ? v$.form.passwordConfirm.$errors[0].$message:''"
    />
    
    <div class="d-flex justify-center mt-3">
      <va-button class="my-0" type="submit"  :disabled="v$.form.$invalid" >Signup</va-button>
    </div>
  </form>
</template>

<script lang="ts">
  import { ref, computed } from 'vue'
  import { useRouter } from 'vue-router'
  import useVuelidate from '@vuelidate/core'
  import { required, email, minLength, maxLength, sameAs, helpers } from '@vuelidate/validators'
  import { first } from '@amcharts/amcharts5/.internal/core/util/Array';
  import axios from 'axios';

  const nameValidator = helpers.regex(/^[a-zA-Z]+[a-zA-Z\s']*$/);

  export default{
    name:"Register",
    setup(){      
      return{       
        v$:useVuelidate()        
      }      
    },
    data(){
      return{
        form:{
          firstName:'',
          lastName:'',
          email:'',
          password:'',
          passwordConfirm:'',
        }
      }
    },

    validations(){
      return {        
        form: {
            firstName:{ 
              required, max:maxLength(20), 
              name_validation:{
                $validator: nameValidator,
                $message: 'Invalid Name, valid name only contain letters, dashes(-), and spaces'
              },
            },
            lastName:{ 
              required, max:maxLength(20), 
              name_validation:{
                $validator: nameValidator,
                $message: 'Invalid Name, valid name only contain letters, dashes(-), and spaces'
              },
            },
            email:{ 
              required, email
            },
            password:{ 
              required, min:minLength(6)
            },
            passwordConfirm:{ 
              required,sameAs: sameAs(this.form.password)
            },
        }
      }
    },
    methods:{
      submit(){
        axios.post('http://127.0.0.1:5000/api/registe',{
          first_name:this.form.firstName,
          last_name:this.form.lastName,
          email:this.form.email,
          password:this.form.password,
          password_confirm:this.form.passwordConfirm,
        }).catch(function (error) {
          if (error.response) {
            // The request was made and the server responded with a status code
            // that falls out of the range of 2xx
            console.log(error.response.data);
          } else if (error.request) {
            // The request was made but no response was received
            // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
            // http.ClientRequest in node.js
            console.log(error.request);
          } else {
            // Something happened in setting up the request that triggered an Error
            console.log('Error', error.message);
          }
        }).then(res =>{
          console.log("B")
        })
      },
    }
  }
</script>
