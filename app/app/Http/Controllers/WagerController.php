<?php

namespace App\Http\Controllers;

use App\Models\Wager;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Carbon;
use App\Http\Requests\Wager\StoreWagersRequest;
use App\Http\Requests\Wager\GetWagersRequest;
use App\Traits\PaginateTrait;

class WagerController extends Controller
{
    use PaginateTrait;

    /**
     * Display a listing of the resource.
     *
     * @param GetWagersRequest $request
     * @return \Illuminate\Contracts\Foundation\Application|\Illuminate\Contracts\Routing\ResponseFactory|\Illuminate\Http\Response
     */
    public function index(GetWagersRequest $request)
    {
        try {
            $page  = $request->input('page', 1);
            $limit = $request->input('limit', 10);
            $selectColumn  = [
                'id',
                'total_wager_value',
                'odds',
                'selling_percentage',
                'selling_price',
                'current_selling_price',
                'percentage_sold',
                'amount_sold',
                'placed_at',
            ];
            $wagers = DB::table('wagers')->orderBy('id', 'desc')->paginate($limit, $selectColumn, 'page', $page);

            return response($this->handleResponseData($wagers));
        } catch (\Exception $exception) {
            return response($exception);
        }
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
     * @return \Illuminate\Contracts\Foundation\Application|\Illuminate\Contracts\Routing\ResponseFactory|\Illuminate\Http\Response
     * @throws \Exception
     */
    public function store(StoreWagersRequest $request)
    {
        DB::beginTransaction();
        try {
            $totalWagerValue      = $request->input('total_wager_value');
            $oods                = $request->input('odds');
            $sellingPercentage   = $request->input('selling_percentage');
            $sellingPrice        = $request->input('selling_price');
            $currentSellingPrice = $request->input('current_selling_price', 0);
            $percentageSold      = $request->input('percentage_sold', 0);
            $amountSold          = $request->input('amount_sold', 0);
            $now                 = Carbon::now();

            $wager = Wager::create([
                'total_wager_value'     => $totalWagerValue,
                'odds'                  => $oods,
                'selling_percentage'    => $sellingPercentage,
                'selling_price'         => $sellingPrice,
                'current_selling_price' => $currentSellingPrice,
                'percentage_sold'       => $percentageSold,
                'amount_sold'           => $amountSold,
                'placed_at'             => $now,
                'created_at'            => $now,
                'updated_at'            => $now
            ]);
            $wager->save();
            DB::commit();


            return response([
                'id'                    => $wager->id,
                'total_wager_value'     => $wager->total_wager_value,
                'odds'                  => $wager->odds,
                'selling_percentage'    => $wager->selling_percentage,
                'selling_price'         => $wager->selling_price,
                'current_selling_price' => $wager->current_selling_price,
                'percentage_sold'       => $wager->percentage_sold,
                'amount_sold'           => $wager->amount_sold,
                'placed_at'             => $wager->placed_at,
            ]);
        } catch (\Exception $exception) {
            DB::rollBack();
            throw $exception;
        }
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Wager  $wager
     * @return \Illuminate\Http\Response
     */
    public function show(Wager $wager)
    {

    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Wager  $wager
     * @return \Illuminate\Http\Response
     */
    public function edit(Wager $wager)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Wager  $wager
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Wager $wager)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Wager  $wager
     * @return \Illuminate\Http\Response
     */
    public function destroy(Wager $wager)
    {
        //
    }
}
