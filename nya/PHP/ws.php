<?php

defined('MAIN_PATH') or define('MAIN_PATH', __DIR__.'/');

// 简单的自动加载
// app\ -> app/ , lib\ -> lib/
spl_autoload_register( function($class){
    $prefix = substr($class, 0, 3);
    if ($prefix === 'app' || $prefix === 'lib') {
        $file = MAIN_PATH.$prefix.'/'.str_replace('\\', '/', substr($class, 4)).'.php';
    } else {
        return;
    }
    if (file_exists($file)) {
        require $file;
    }
} );

new \lib\Worker(2333);