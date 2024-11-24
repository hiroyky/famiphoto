<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>{{ $t('applicationSettings') }}</h1>

        <v-container>
          <v-row>
            <v-col md="8">
              <v-card>
                <v-card-title>インデックス更新</v-card-title>
                <v-card-subtitle>インデックス更新について</v-card-subtitle>
                <v-card-text>
                  <p>写真ディレクトリ以下の画像ファイル(.jpg, .raw, .arw)を毎日定時にインデックス更新します。
                  写真ディレクトリ配下に直接画像ファイルを設置した場合は、インデックス更新で追加されます。</p>
                  <p><nuxt-link to="/upload">写真をアップロード</nuxt-link>からアップロードした写真は即時インデックス登録されます。</p>
                </v-card-text>

                <v-list>
                  <v-list-subheader>インデックス提示自動更新を実行する時間</v-list-subheader>
                  <v-list-item>
                    <v-list-item-action>
                      <v-text-field label="毎日実行する日時" class="pr-4"></v-text-field>
                      <v-btn>反映</v-btn>
                    </v-list-item-action>
                  </v-list-item>

                  <v-list-item>
                    <template #subtitle>前回実行日時:</template>
                    2024年6月27日2時37分
                  </v-list-item>

                  <v-list-item>
                    インデックス更新は定時実行されますが、今すぐ実行することもできます。
                  </v-list-item>
                  <v-list-item>
                    <v-btn variant="flat" color="primary">今すぐインデックスを更新する</v-btn>
                    <template #append>
                      <div>
                        <v-progress-circular indeterminate></v-progress-circular>
                        <span class="pl-1">現在、インデックス更新を処理中です..</span>
                      </div>
                    </template>
                  </v-list-item>

                </v-list>
              </v-card>
            </v-col>
          </v-row>
        </v-container>

      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
</script>

<style lang="scss">

</style>