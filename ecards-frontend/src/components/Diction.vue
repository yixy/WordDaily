<template>
  <div class="diction-container">
    <div class="tabs">
      <button
        :class="{ active: activeTab === 'category' }"
        @click="activeTab = 'category'"
      >
        分类
      </button>
      <button
        :class="{ active: activeTab === 'statistics' }"
        @click="activeTab = 'statistics'"
      >
        统计
      </button>
      <button
        :class="{ active: activeTab === 'import' }"
        @click="activeTab = 'import'"
      >
        导入
      </button>
    </div>
    <div class="content">
      <div v-if="activeTab === 'category'" class="category">
        <div class="tags">
          <span
            v-for="tag in tags"
            :key="tag"
            class="tag"
            @click="selectTag(tag)"
            >{{ tag }}</span
          >
        </div>
        <div class="word-list">
          <div
            v-for="word in words"
            :key="word.word"
            class="word-item"
            @click="toggleMeaning(word)"
          >
            <div class="word">{{ word.word }}</div>
            <div v-if="word.showMeaning" class="meaning">
              {{ word.meaning }}
            </div>
          </div>
        </div>
      </div>
      <div v-if="activeTab === 'statistics'" class="statistics">
        <div class="stat-item">未学习单词: {{ stats.unlearned }}</div>
        <div class="stat-item">7天复习单词数: {{ stats.review7Days }}</div>
        <div class="stat-item">30天复习单词数: {{ stats.review30Days }}</div>
        <div class="stat-item">100天复习单词数: {{ stats.review100Days }}</div>
        <div class="stat-item">单词总数: {{ stats.total }}</div>
      </div>
      <div v-if="activeTab === 'import'" class="import">
        <button class="import-button" @click="importWords">批量导入</button>
        <textarea
          v-model="importText"
          placeholder="输入单词和标签..."
        ></textarea>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeTab: "category",
      tags: ["自然景观", "手机问题"],
      words: [
        {
          word: "Scenery",
          meaning:
            "The natural features of a landscape considered in terms of their appearance, especially when picturesque.",
          showMeaning: false,
        },
        {
          word: "Forest",
          meaning: "A large area covered chiefly with trees and undergrowth.",
          showMeaning: false,
        },
      ],
      stats: {
        unlearned: 10,
        review7Days: 5,
        review30Days: 20,
        review100Days: 50,
        total: 100,
      },
      importText: "",
    };
  },
  methods: {
    selectTag(tag) {
      // 选择标签的逻辑
    },
    toggleMeaning(word) {
      word.showMeaning = !word.showMeaning;
    },
    async importWords() {
      try {
        const response = await this.$http.post("/api/import", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ text: this.importText }),
        });
        if (response.ok) {
          alert("导入成功");
        } else {
          alert("导入失败");
        }
      } catch (error) {
        alert(error.message);
      }
    },
  },
};
</script>

<style scoped>
.diction-container {
  padding: 20px;
}

.tabs {
  display: flex;
  justify-content: space-around;
  margin-bottom: 20px;
}

.tabs button {
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.tabs button.active {
  background-color: #007bff;
  color: white;
}

.category .tags {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.category .tag {
  padding: 5px 10px;
  background-color: #e9ecef;
  border-radius: 5px;
  cursor: pointer;
}

.word-list {
  margin-top: 20px;
}

.word-item {
  padding: 10px;
  border-bottom: 1px solid #ddd;
  cursor: pointer;
}

.word-item .word {
  font-weight: bold;
}

.word-item .meaning {
  margin-top: 5px;
  color: #666;
}

.statistics .stat-item {
  margin-bottom: 10px;
}

.import .import-button {
  margin-bottom: 10px;
  padding: 10px 20px;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.import textarea {
  width: 100%;
  height: 150px;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}
</style>
