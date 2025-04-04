<template>
  <div class="punch-container">
    <div class="card" @click="flipCard">
      <div class="card-inner" :class="{ 'is-flipped': isFlipped }">
        <div class="card-front">
          <h2>{{ currentWord.word }}</h2>
        </div>
        <div class="card-back">
          <p>{{ currentWord.meaning }}</p>
          <p class="example">{{ currentWord.example }}</p>
        </div>
      </div>
    </div>
    <div class="buttons">
      <button class="retain" @click="retainWord">保留</button>
      <button class="break" @click="breakWord">击破</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isFlipped: false,
      currentWord: {
        word: 'Scenery',
        meaning: 'The natural features of a landscape considered in terms of their appearance, especially when picturesque.',
        example: 'The scenery in the mountains was breathtaking.'
      }
    };
  },
  methods: {
    flipCard() {
      this.isFlipped = !this.isFlipped;
    },
    retainWord() {
      // 保留单词的逻辑
    },
    breakWord() {
      // 击破单词的逻辑
    }
  }
};
</script>

<style scoped>
.punch-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.card {
  width: 300px;
  height: 200px;
  perspective: 1000px;
}

.card-inner {
  position: relative;
  width: 100%;
  height: 100%;
  text-align: center;
  transition: transform 0.6s;
  transform-style: preserve-3d;
}

.card-inner.is-flipped {
  transform: rotateY(180deg);
}

.card-front, .card-back {
  position: absolute;
  width: 100%;
  height: 100%;
  backface-visibility: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.card-front {
  background-color: #f8f9fa;
}

.card-back {
  background-color: #e9ecef;
  transform: rotateY(180deg);
}

.buttons {
  margin-top: 20px;
}

button {
  padding: 10px 20px;
  margin: 0 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.retain {
  background-color: #28a745;
  color: white;
}

.break {
  background-color: #dc3545;
  color: white;
}
</style>