/*
* 浏览器存储
* */

export const SDB = new (class SimpleDB {
    set = (key, value) => localStorage.setItem(key, value)
    get = (key) => localStorage.getItem(key)
    remove = (key) => localStorage.removeItem(key)
})