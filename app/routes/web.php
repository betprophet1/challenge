<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\WagerController;
use App\Http\Controllers\PurchaseController;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});

Route::post('/wagers', [WagerController::class, 'store']);
Route::get('/wagers', [WagerController::class, 'index']);
Route::post('/buy/{wager_id}', [PurchaseController::class, 'buy']);
