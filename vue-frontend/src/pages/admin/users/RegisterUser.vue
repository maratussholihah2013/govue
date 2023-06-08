<template>
  <div class="form-elements">
    <div class="row">
      <div class="flex xs12">
        <va-card>
          <va-card-title>Register User Baru</va-card-title>
          <va-card-content>
            <form @submit="submit">
              <div class="row">
                <div class="flex md4 sm6 xs12">
                  <va-input v-model="firstName" placeholder="First Name" />
                </div>
                <div class="flex md4 sm6 xs12">
                  <va-input v-model="lastName" placeholder="Last Name" />
                </div>
                <div class="flex md4 sm6 xs12">
                  <va-input v-model="email" type="email" placeholder="Email">
                    <template #prepend>
                      <va-icon color="gray" name="envelope" />
                    </template>
                  </va-input>
                </div>
                <div class="flex md4 sm6 xs12">
                  <va-input v-model="password" type="password" placeholder="Password" />
                </div>
                <div class="flex md4 sm6 xs12">
                  <va-input v-model="passwordConfirm" type="password" placeholder="Konfirmasi Password" />
                </div>
                <div class="flex md4 sm6 xs12">
                  <va-input
                    v-model="address"
                    placeholder="Alamat"
                    messages="Alamat detail dengan RT/RW "
                  />
                </div>
                <div class="flex md4 sm6 xs12">
                  <!--
                    manual-input : typeable date
                    disabled : disabled changeable
                    mode="multiple" : multi select date
                    mode="range" : range date
                    clearable : button clear select date
                    first-weekday="Monday" : first date display per row is Monday                    
                    highlight-weekend : make saturday and sunday red colored
                    label="LABEL" : judul di atas input
                    <va-time-input> : untuk memilih jam
                  -->
                  <va-date-input v-model="birthdate" placeholder="Pilih Tanggal Lahir"/>
                </div>
                <div class="flex md6 xs12">
                  <va-select
                    v-model="placebirth"
                    label="Tanggal Lahir"
                    text-by="description"
                    track-by="id"
                    multiple
                    searchable
                    :options="kotas"
                  />
                </div>
                <div class="flex md4 sm6 xs12">
                  <va-switch v-model="status" size="small" label="Small toggle" class="mr-4 mb-2" />
                </div>
                <div class="flex offset-md10">
                  <va-button class="w-100 mr-2 mb-2" type="submit"> Submit</va-button>
                </div>
              </div>
            </form>
          </va-card-content>
          <va-divider />
          <va-card-content class="row">
          </va-card-content>
        </va-card>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
  import { ref } from 'vue'
  //import { useI18n } from 'vue-i18n'
  //const { t } = useI18n()

  export default{
    name:"RegisterUser",
    setup(){
      const firstName  = ref('');
      const lastName = ref('');
      const email = ref('');
      const password = ref('');
      const passwordConfirm = ref('');

      const submit = () => {
        console.log({
          first_name: firstName.value,
          last_name: lastName.value,
          email: email.value,
          password: password.value,
          password_confirm: passwordConfirm.value,
        })
      }

      return {
        firstName,
        lastName,
        email,
        password,
        passwordConfirm,
        submit
      }
    }
  }

</script>

<style lang="scss" scoped>
  fieldset {
    margin-bottom: 0.5rem;
  }
</style>
