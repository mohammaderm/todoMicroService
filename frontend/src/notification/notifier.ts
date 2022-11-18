// =============== Libraries =============== //
import { Store } from 'react-notifications-component';

type NotifType = "default" | "success" | "danger" | "warning" | "info"

const createNotification = (code: number, message: string, title: string, time: number) => {
    let type: NotifType = "info"
    if (code > 199 && code < 300) {
        type = "success"
        if (title === "") {
            title = "Success"
        }
    } else if (code > 399 && code < 500) {
        type = "warning"
        if (title === "") {
            title = "Bad Request"
        }
    } else if (code > 499 && code < 600) {
        type = "danger"
        if (title === "") {
            title = "Internal Server Error"
        }
    }

    if (time === 0) {
        time = 5000
    }

    Store.addNotification({
        title: title,
        message: message,
        type: type,
        insert: "top",
        container: "top-center",
        animationIn: ["animate__animated", "animate__fadeIn"],
        animationOut: ["animate__animated", "animate__fadeOut"],
        dismiss: {
            duration: time,
            onScreen: true
        }
    });
}

export default createNotification;