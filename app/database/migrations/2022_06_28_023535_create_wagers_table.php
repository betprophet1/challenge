<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('wagers', function (Blueprint $table) {
            $table->increments('id');
            $table->float('total_wager_value');
            $table->float('odds');
            $table->float('selling_percentage');
            $table->float('selling_price');
            $table->float('current_selling_price')->default(0);
            $table->float('percentage_sold')->default(0);
            $table->float('amount_sold')->default(0);
            $table->timestamp('placed_at');
            $table->timestamps();
            $table->timestamp('deleted_at')->nullable();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('wagers');
    }
};
