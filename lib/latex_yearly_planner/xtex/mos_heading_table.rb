# frozen_string_literal: true

module LatexYearlyPlanner
  module XTeX
    class MosHeadingTable
      attr_reader :page_name, :placement, :tabularx, :navigation

      def initialize(page_name:, placement:, tabularx:, navigation:)
        @page_name = page_name
        @placement = placement
        @tabularx = tabularx
        @navigation = navigation
      end

      def to_s
        table = TeX::TabularX.new(**tabularx)
        table.formatting = TeX::TableFormatting.new(layout)
        table.add_row(TeX::TableRow.new(cell_mix.flatten))

        table.to_s
      end

      private

      def layout
        placement.map { |item| item.position || "|#{nav_layout_part}|" }.join
      end

      def nav_layout_part
        navigation.size.times.map { 'l' }.split.join('|')
      end

      def cell_mix
        @cell_mix ||= placement.map { |item| method(item.function).call }
      end

      def empty_cell_filler
        TeX::TableCell.new
      end
    end
  end
end
