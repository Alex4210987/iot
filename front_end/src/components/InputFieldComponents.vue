<script>
export default {
  props: ['type', 'label'],
  data() {
    return {
      isActive: false,
      value: '',
    };
  },
  methods: {
    updateValue($event) {
      this.$emit('inputEvent', $event.target.value);
    }
  },
}
</script>

<template>
  <div class="input-field">
    <input
        :type="type"
        @input="updateValue"
        v-model="value"
        class="input"
        @focus="isActive = true"
        @blur="isActive = false"
    />
    <label class="label"  :class="{ 'active': isActive || value.length > 0, 'inactive': !isActive && value.length === 0 }">{{label}}</label>
  </div>
</template>

<style scoped>
.input-field {
  width: 100%;
  height: 10%;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.input{
  background: transparent;
  border: none;
  border-bottom: 2px solid rgba(0, 0, 0, 0.5);
  outline: none;
  color: rgba(0, 0, 0, 0.5);
  font-size: 0.8rem;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 50%;
  height: 30px;
}

.label {
  transition: all 0.3s ease;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 50%;
  height: 20px;
  pointer-events: none;
  color: rgba(0, 0, 0, 0.5);
  font-size: 0.8rem;
}

.label.active {
  transform: translate(-50%, calc(-50% - 20px));
}

.label.inactive {
  transform: translate(-50%, -50%);
}
</style>
