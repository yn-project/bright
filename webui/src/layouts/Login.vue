<template>
  <a-space class="top-title" :size="20">
    <div style="display: inline-block;width: 20px;"></div>
    <a-image :src="logo" :width="64" />
    <h1 class="main-text-2">农产品区块链平台</h1>
  </a-space>
  <a-card title="登 录" style="width: 400px;margin: auto;">

    <a-form :model="formState" name="normal_login" class="login-form" @finish="onFinish" @finishFailed="onFinishFailed">
      <a-form-item label="Username" name="用户名" :rules="[{ required: true, message: 'Please input your username!' }]">
        <a-input v-model:value="formState.username">
          <template #prefix>
            <UserOutlined class="site-form-item-icon" />
          </template>
        </a-input>
      </a-form-item>

      <a-form-item label="Password" name="密 码" :rules="[{ required: true, message: 'Please input your password!' }]">
        <a-input-password v-model:value="formState.password">
          <template #prefix>
            <LockOutlined class="site-form-item-icon" />
          </template>
        </a-input-password>
      </a-form-item>

      <div class="login-form-wrap">
        <a-form-item name="remember" no-style>
          <a-checkbox v-model:checked="formState.remember">Remember me</a-checkbox>
        </a-form-item>
        <a class="login-form-forgot" href="">Forgot password</a>
      </div>

      <a-form-item>
        <div style="display: grid;">
          <a-button :disabled="disabled" type="primary" html-type="submit" class="login-form-button"
            style="display: inline-block;justify-self: flex-end;align-self:flex-end;">
            Log in
          </a-button>
        </div>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<script lang="ts">
import { defineComponent, reactive, computed } from 'vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';
import logo from '../assets/logo.svg';

interface FormState {
  username: string;
  password: string;
  remember: boolean;
}
export default defineComponent({
  components: {
    UserOutlined,
    LockOutlined,
  },
  setup() {
    const formState = reactive<FormState>({
      username: '',
      password: '',
      remember: true,
    });
    const onFinish = (values: any) => {
      console.log('Success:', values);
    };

    const onFinishFailed = (errorInfo: any) => {
      console.log('Failed:', errorInfo);
    };
    const disabled = computed(() => {
      return !(formState.username && formState.password);
    });
    return {
      logo,
      formState,
      onFinish,
      onFinishFailed,
      disabled,
    };
  },
});
</script>
<style>
#components-form-demo-normal-login .login-form {
  max-width: 300px;
}

#components-form-demo-normal-login .login-form-wrap {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

#components-form-demo-normal-login .login-form-forgot {
  margin-bottom: 24px;
}

#components-form-demo-normal-login .login-form-button {
  width: 100%;
}
</style>
<style scoped>
h1 {
  font-weight: bold;
  line-height: 60px;
}

.top-title {
  padding: 30px;
}
</style>
