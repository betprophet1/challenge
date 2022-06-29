<?php

namespace App\Http\Controllers;

use App\Models\Purchase;
use App\Models\Wager;
use Illuminate\Http\Request;
use Illuminate\Support\Carbon;
use Illuminate\Support\Facades\DB;
use App\Http\Requests\Purchase\BuyPurchasesRequest;

class PurchaseController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        //
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Purchase  $purchase
     * @return \Illuminate\Http\Response
     */
    public function show(Purchase $purchase)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Purchase  $purchase
     * @return \Illuminate\Http\Response
     */
    public function edit(Purchase $purchase)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Purchase  $purchase
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Purchase $purchase)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Purchase  $purchase
     * @return \Illuminate\Http\Response
     */
    public function destroy(Purchase $purchase)
    {
        //
    }

    /**
     * @param BuyPurchasesRequest $request
     * @param $wagerId
     * @return \Illuminate\Contracts\Foundation\Application|\Illuminate\Contracts\Routing\ResponseFactory|\Illuminate\Http\Response
     * @throws \Exception
     */

    public function buy(BuyPurchasesRequest $request, $wagerId)
    {
        DB::beginTransaction();
        try {
            $buyingPrice = $request->input('buying_price');
            $purchase = Purchase::create([
                'wager_id'     => $wagerId,
                'buying_price' => $buyingPrice,
                'bought_at'    => Carbon::now(),
            ]);

            $purchase->save();
            DB::commit();

            $countPurchase = DB::table('purchases')->where('wager_id', $wagerId)->count();

            $oldWager = DB::table('wagers')->find($wagerId);
            $wager = [
                'current_selling_price' => $purchase->buying_price,
                'percentage_sold'       => $purchase->buying_price / $oldWager->total_wager_value * 100,
                'amount_sold'           => $countPurchase,
            ];

            Wager::whereId($wagerId)->update($wager);

            return response([
                'id'           => $purchase->id,
                'wager_id'     => $purchase->wager_id,
                'buying_price' => $purchase->buying_price,
                'bought_at'    => $purchase->bought_at,
            ]);

        } catch (\Exception $exception) {
            DB::rollBack();
            throw $exception;
        }
    }
}
