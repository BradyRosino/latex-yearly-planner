# frozen_string_literal: true

module LatexYearlyPlanner
  module XTeX
    class CalendarLarge
      attr_reader :month,
                  :show_week_numbers,
                  :week_number_placement,
                  :cell_height,
                  :vertical_stretch,
                  :horizontal_spacing,
                  :horizontal_lines,
                  :name_cell_height

      def initialize(month, **options)
        @month = month

        @show_week_numbers = options.fetch(:show_week_numbers, true)
        @week_number_placement = options.fetch(:week_number_placement, :left).downcase.to_sym
        @cell_height = options.fetch(:cell_height, '2cm')
        @vertical_stretch = options.fetch(:vertical_stretch, 1.0)
        @horizontal_spacing = options.fetch(:horizontal_spacing, nil)
        @horizontal_lines = options.fetch(:horizontal_lines, true)
        @name_cell_height = options.fetch(:name_cell_height, '5mm')
      end

      def to_s
        table = TeX::TabularX.new(**table_options)

        table.add_row(header)

        week_rows.each { |row| table.add_row(row) }

        table
      end

      private

      def table_options
        { format:, horizontal_lines:, vertical_stretch:, horizontal_spacing: }.compact
      end

      def format
        space = '@{\hspace{1.5mm}}'

        return "#{"|#{space}X" * 7}|" unless show_week_numbers

        return "|c|#{"#{space}X|" * 7}" if week_number_placement == :left

        "#{"|#{space}X" * 7}|c|"
      end

      def header
        row = month.weekdays_short.map { |day| "\\hfil{}#{day}" }
        row[0] = "#{MinHeight.new(name_cell_height)}#{row[0]}"

        return row unless show_week_numbers

        return ['W'] + row if week_number_placement == :left

        row + ['W']
      end

      def week_rows
        month.weeks.map do |week|
          row = week.days.map { |day| day ? "{#{day.mday}}" : '' }
          row[0] = "\\adjustbox{valign=t}{\\vphantom{\\rule{0pt}{#{cell_height}}}}#{row[0]}"

          next row unless show_week_numbers

          row.unshift(rotated(week)) if week_number_placement == :left
          row.push(rotated(week)) if week_number_placement == :right

          row
        end
      end

      def rotated(week)
        "\\rotatebox[origin=tr]{90}{\\parbox{#{cell_height}}{\\hfil{}#{week_name_box(week)}}}"
      end

      def week_name_box(week)
        "#{XTeX::MinHeight.new('6mm')}Week #{week.number}"
      end
    end
  end
end
