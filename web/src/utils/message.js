import { Message, MessageBox } from "element-ui";

export function openMessage(icon = '', title, message, cancelButtonText, confirmButtonText, showConfirmButton = true, showCancelButton = true) {
  return new Promise((resolve, reject) => {
    MessageBox({
      customClass: 'wk-messagebox',
      iconClass: [icon],
      center: true,
      title,
      message,
      cancelButtonText: cancelButtonText || '取消',
      confirmButtonText: confirmButtonText || '确定',
      cancelButtonClass: 'wk-default-button',
      confirmButtonClass: 'wk-primary-button',
      showConfirmButton: showConfirmButton,
      showCancelButton: showCancelButton,
      callback: resolve
    })
  })
}

export class WKMessage {
  static wkOptions = {
    customClass: "wk-message",
    center: true,
    duration: 3000,
    showClose: true
  }
  static success (msg, options) {
    WKMessage.wkOptions.message = msg;
    WKMessage.message('success', WKMessage.wkOptions)
  }
 
  static warning (msg, options) {
    WKMessage.wkOptions.message = msg;
    WKMessage.message('warning', WKMessage.wkOptions)
  }
 
  static info (msg, options) {
    WKMessage.wkOptions.message = msg;
    WKMessage.message('info', WKMessage.wkOptions)
  }
 
  static error (msg, options) {
    WKMessage.wkOptions.message = msg;
    WKMessage.message('error', WKMessage.wkOptions)
  }
 
  static message (type, options) {
    Message[type](options)
    // 如果需要数量 请自行限制
    // const messageDom = document.getElementsByClassName('el-message')[0]
    // if (messageDom === undefined) {
    //   Message[type](options)
    // }
  }
}