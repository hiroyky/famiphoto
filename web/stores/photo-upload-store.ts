
interface State {
    queue: FileUploadQueueItem[]
    progressing: boolean
}

type UploadStatus = 'standby' | 'progressing'| 'completed' | 'failed'

export interface FileUploadQueueItem {
    file: File
    status: UploadStatus
}

const parallelNum = 2

export const usePhotoUploadStore = defineStore('photoUpload', {
    state: (): State => ({
        queue: [],
        progressing: false,
    }),

    actions: {
        enqueueToUploadFiles(files: File[]) {
            const fileItems: FileUploadQueueItem[] = files.map(f => ({
                file: f,
                status: 'standby',
            }))
            this.queue.push(...fileItems)
            this._uploadQueuing()
        },
        clear() {
            this.queue = []
        },
        async _uploadQueuing() {
            const standbyFiles = filterStandbyFiles(this.queue)
            if (standbyFiles.length===0) {
                this.progressing = false
                return
            }

            this.progressing = true
            console.log('progresing', this.progressing)
            const { client } = useGqlStore()
            const apiStore = useApiStore()

            await Promise.all(standbyFiles.slice(0, parallelNum).map(async (i) => {
                try {
                    i.status='progressing'
                    const res = await client.uploadPhoto()
                    await apiStore.uploadFile(i.file, res.uploadPhoto.uploadUrl)
                    i.status = 'completed'
                } catch(err) {
                    i.status = 'failed'
                    console.error(err)
                }
            }))

            this._uploadQueuing()
        },
    },
})

function filterStandbyFiles(queue: FileUploadQueueItem[]): FileUploadQueueItem[] {
    return queue.filter(m => m.status === 'standby')
}