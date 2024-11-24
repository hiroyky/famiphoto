import {useDisplay} from "vuetify";

export function useNavigation() {
    const { mobile } = useDisplay()
    const display = useState<boolean>('navigation-display', () => true)
    function toggle() {
        display.value = !display.value
    }
    function init() {
        display.value = !mobile.value
    }

    return {
        init,
        display,
        toggle,
        temporary: () => mobile.value
    }
}