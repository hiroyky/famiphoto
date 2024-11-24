import {defineStore} from "pinia";

interface State {
    alertQueue: {
        body: string,
        isPermanent?: boolean,
        type: 'fatal' | 'error' | 'warn' | 'message'
    }[]
}

const useAlertStore = defineStore('alert', {
    state: (): State => ({
        alertQueue: [],
    }),

    actions: {
        displayFatal (body: string, isPermanent = false) {
            this.alertQueue.push({ body, isPermanent, type: 'fatal' })
        },
        displayError (body: string, isPermanent = false) {
            this.alertQueue.push({ body, isPermanent, type: 'error' })
        },
        displayWarn (body: string, isPermanent = false) {
            this.alertQueue.push({ body, isPermanent, type: 'warn' })
        },
        displayMessage (body: string, isPermanent = false) {
            this.alertQueue.push({ body, isPermanent, type: 'message' })
        },
        deque () {
            return this.alertQueue.shift()
        },
    },
})

export default useAlertStore